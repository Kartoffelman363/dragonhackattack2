import React, {
    ReactElement,
    useCallback,
    useEffect,
    useRef,
    useState,
} from 'react';
import { Subscription } from 'rxjs';

import './Chatbot.scss';
import JointPlusService from '../../services/joint-plus.service';
import JsonEditor from './JsonEditor/JsonEditor';
import Inspector from './Inspector/Inspector';
import EventBusServiceContext from '../../services/event-bus-service.context';
import { EventBusService } from '../../services/event-bus.service';
import { SharedEvents } from '../../joint-plus/controller';
import {
    importGraphFromJSON,
    loadStencilShapes,
    zoomToFit,
} from '../../joint-plus/actions';
import { STENCIL_WIDTH } from '../../theme';

import exampleGraphJSON from '../../joint-plus/config/example-graph.json';

const Chatbot = (): ReactElement => {
    const elementRef = useRef(null);
    const toolbarRef = useRef(null);
    const stencilRef = useRef(null);
    const paperRef = useRef(null);
    const jsonForBackend = useRef(null);

    const [joint, setJoint] = useState(null);
    const [eventBusService] = useState(new EventBusService());
    const [stencilOpened, setStencilOpened] = useState(true);
    const [jsonEditorOpened, setJsonEditorOpened] = useState(true);
    const [fileJSON, setFileJSON] = useState(null);
    const [subscriptions] = useState(new Subscription());

    const openFile = useCallback(
        (json: Object): void => {
            setFileJSON(json);
            importGraphFromJSON(joint, json);
            zoomToFit(joint);
        },
        [joint],
    );

    const onStart = useCallback((): void => {
        loadStencilShapes(joint);
        openFile(exampleGraphJSON);
    }, [joint, openFile]);

    const onJsonEditorChange = useCallback(
        (json: Object): void => {
            if (joint) {
                importGraphFromJSON(joint, json);
            }
        },
        [joint],
    );

    const onJointGraphChange = useCallback((json: Object): void => {
        jsonForBackend.current = json;
    }, []);

    const onStencilToggle = useCallback((): void => {
        if (!joint) {
            return;
        }
        const { scroller, stencil } = joint;
        if (stencilOpened) {
            stencil.unfreeze();
            scroller.el.scrollLeft += STENCIL_WIDTH;
        } else {
            stencil.freeze();
            scroller.el.scrollLeft -= STENCIL_WIDTH;
        }
    }, [joint, stencilOpened]);

    const toggleJsonEditor = (): void => {
        setJsonEditorOpened(!jsonEditorOpened);
    };

    const sendJsonToBackend = async (): Promise<void> => {
        console.log('JSON by Balenciaga:\n', jsonForBackend.current);
        class InputVariable {
            VarName: string;
            Id: string;
            Type: string;
            Value: string;
        }
        class OutputVariable {
            VarName: string;
            Id: string;
            Type: string;
        }
        class Block {
            Id: string;
            InputVariables: InputVariable[]; // array of class InputVariable where VarName is same as input parameter of Code: string
            // and Value is UUID of some OutputVariable
            OutputVariables: OutputVariable[]; // array of class InputVariable where VarName is
            Code: string; // example:     myFunction(varName1, varName2)
            // Block.Blocks is NO MORE!!!!
            //Blocks: Block[];
        }
        class ResJson {
            Metadata: string;
            InitialVariables: InputVariable[]; // array of class InputVariable where VarName is UUID and Value is external value defined at start
            Blocks: Block[]; // array of class Blocks
        }

        // Tale JSON ti dobiš
        let resJson: ResJson = {
            Metadata: JSON.stringify(jsonForBackend.current),
            InitialVariables: [] as InputVariable[], // array of class InputVariable where VarName is UUID and Value is external value defined at start
            Blocks: [] as Block[], // array of class Blocks
        } as ResJson;

        if (!jsonForBackend.current.cells) {
            console.log('Invalid jsonForBackend.');
            return;
        }

        const cells: Array<any> = jsonForBackend.current.cells;
        const flowchartStart: any = cells.find((cell) => {
            return cell.type == 'app.FlowchartStart';
        });
        console.log('flowchartStart: ', flowchartStart);
        cells.forEach((cell) => {
            if (cell.type != 'app.Message') return;
            let block: Block = {} as Block;
            block.Id = cell.id;
            block.Code = cell.function;
            block.InputVariables = [] as InputVariable[];
            block.OutputVariables = [] as OutputVariable[];
            cell.ports.items.forEach((port: any) => {
                if (port.group == 'in') {
                    block.InputVariables.push({
                        VarName: port.attrs.portLabel.text,
                        Id: port.id, //Ignore on backend
                        Type: port.type,
                        Value: '', //UUID of an OutputVariable ( OutputVariable.VarName )
                    } as InputVariable);
                } else if (port.group == 'out') {
                    block.OutputVariables.push({
                        VarName: port.id, //UUID
                        Id: port.id, //Ignore on backend
                        Type: port.type,
                    } as OutputVariable);
                }
            });

            resJson.Blocks.push(block);
        });

        cells.forEach((cell) => {
            if (cell.type != 'app.Link') return;
            const source = cell.source.port;
            const target = cell.target.port;
            resJson.Blocks.forEach((_, blockIdx: number) => {
                resJson.Blocks[blockIdx].InputVariables.forEach(
                    (inputVar: InputVariable, inputVarIdx: number) => {
                        if (inputVar.Id == target) {
                            resJson.Blocks[blockIdx].InputVariables[
                                inputVarIdx
                            ].Value = source;
                            if (flowchartStart.ports.items[0].id == source) {
                                resJson.InitialVariables.push({
                                    VarName: inputVar.VarName,
                                    Id: inputVar.Value,
                                    Type: inputVar.Type,
                                    Value: '',
                                } as InputVariable);
                            }
                        }
                    },
                );
                return;
            });
        });

        try {
            await fetch('http://127.0.0.1:8000/workflows', {
                method: 'POST',
                body: JSON.stringify(resJson),
                headers: {
                    'Content-type': 'application/json; charset=UTF-8',
                },
            });
        } catch {
            console.log('Failed to send POST workflows.');
        }
        console.log('JSON by H&M:\n', resJson, JSON.stringify(resJson));
    };

    const toggleStencil = (): void => {
        setStencilOpened(!stencilOpened);
    };

    useEffect((): void => {
        onStencilToggle();
    }, [stencilOpened, onStencilToggle]);

    const setStencilContainerSize = useCallback((): void => {
        stencilRef.current.style.width = `${STENCIL_WIDTH}px`;
    }, []);

    useEffect(() => {
        subscriptions.add(
            eventBusService.subscribe(
                SharedEvents.GRAPH_CHANGED,
                (json: Object) => onJointGraphChange(json),
            ),
        );
        subscriptions.add(
            eventBusService.subscribe(
                SharedEvents.JSON_EDITOR_CHANGED,
                (json: Object) => onJsonEditorChange(json),
            ),
        );
    }, [
        eventBusService,
        subscriptions,
        onJointGraphChange,
        onJsonEditorChange,
    ]);

    useEffect(() => {
        setJoint(
            new JointPlusService(
                elementRef.current,
                paperRef.current,
                stencilRef.current,
                toolbarRef.current,
                eventBusService,
            ),
        );
    }, [eventBusService]);

    useEffect(() => {
        if (!joint) {
            return;
        }
        setStencilContainerSize();
        onStart();
    }, [joint, onStart, setStencilContainerSize]);

    useEffect(() => {
        if (!joint) {
            return;
        }
        return () => {
            subscriptions.unsubscribe();
            joint.destroy();
        };
    }, [joint, subscriptions]);

    return (
        <EventBusServiceContext.Provider value={eventBusService}>
            <div ref={elementRef} className="joint-scope chatbot">
                <div ref={toolbarRef} />
                <div className="side-bar">
                    <div className="toggle-bar">
                        <div
                            onClick={toggleStencil}
                            className={
                                'icon toggle-stencil ' +
                                (!stencilOpened ? 'disabled-icon' : '')
                            }
                            data-tooltip="Toggle Element Palette"
                            data-tooltip-position-selector=".toggle-bar"
                        />
                        <div
                            onClick={toggleJsonEditor}
                            className={
                                'icon toggle-editor ' +
                                (!jsonEditorOpened ? 'disabled-icon' : '')
                            }
                            data-tooltip="Toggle JSON Editor"
                            data-tooltip-position-selector=".toggle-bar"
                        />
                        <div
                            onClick={sendJsonToBackend}
                            className="icon toggle-editor "
                            data-tooltip="Send JSON"
                            data-tooltip-position-selector=".toggle-bar"
                        />
                    </div>
                    <div
                        ref={stencilRef}
                        style={{ display: stencilOpened ? 'initial' : 'none' }}
                        className="stencil-container"
                    />
                </div>
                <div className="main-container">
                    <div ref={paperRef} className="paper-container" />
                    <div
                        style={{
                            display: jsonEditorOpened ? 'initial' : 'none',
                        }}
                    >
                        <JsonEditor content={fileJSON} />
                    </div>
                </div>
                <Inspector />
            </div>
        </EventBusServiceContext.Provider>
    );
};

export default Chatbot;

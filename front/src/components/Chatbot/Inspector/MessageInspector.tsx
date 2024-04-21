import React, {
    ChangeEvent,
    ReactElement,
    useCallback,
    useEffect,
    useState,
} from 'react';
import { shapes } from '@joint/plus';

import { useBaseInspector } from './useBaseInspector';
import Input from '../Input/Input';

interface Props {
    cell: shapes.app.Message;
}

interface InspectorPort {
    id: string;
    label: string;
}

const cellProps = {
    label: ['attrs', 'label', 'text'],
    description: ['attrs', 'description', 'text'],
    icon: ['attrs', 'icon', 'xlinkHref'],
    blockFunction: ['attrs', 'blockFunction', 'text'],
    portLabel: ['attrs', 'portLabel', 'text'],
};

const MessageInspector = (props: Props): ReactElement => {
    const { cell } = props;

    const [label, setLabel] = useState<string>('');
    const [description, setDescription] = useState<string>('');
    const [icon, setIcon] = useState<string>('');

    const assignFormFields = useCallback((): void => {
        setLabel(cell.prop(cellProps.label));
        setDescription(cell.prop(cellProps.description));
        setIcon(cell.prop(cellProps.icon));
    }, [cell]);

    const changeCellProp = useBaseInspector({ cell, assignFormFields });

    return (
        <>
            <h1>Component</h1>

            <label htmlFor="label">Label</label>
            <Input
                id="label"
                type="text"
                placeholder="Enter label"
                value={label}
                onChange={(e: ChangeEvent<HTMLInputElement>) =>
                    changeCellProp(cellProps.label, e.target.value)
                }
            />
            <label htmlFor="description">Description</label>
            <Input
                id="description"
                type="text"
                placeholder="Enter description"
                value={description}
                onChange={(e: ChangeEvent<HTMLInputElement>) =>
                    changeCellProp(cellProps.description, e.target.value)
                }
            />
            <label htmlFor="icon">Icon (Base64)</label>
            <span className="icon-input" />
            <Input
                id="icon"
                type="text"
                placeholder="Enter icon"
                value={icon}
                spellCheck={false}
                onChange={(e: ChangeEvent<HTMLInputElement>) =>
                    changeCellProp(cellProps.icon, e.target.value)
                }
            />
        </>
    );
};

export default MessageInspector;

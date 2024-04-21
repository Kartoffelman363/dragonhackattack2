/*! JointJS+ v4.0.1 - HTML5 Diagramming Framework - TRIAL VERSION

Copyright (c) 2024 client IO

 2024-04-19 


This Source Code Form is subject to the terms of the JointJS+ Trial License
, v. 2.0. If a copy of the JointJS+ License was not distributed with this
file, You can obtain one at https://www.jointjs.com/license
 or from the JointJS+ archive as was distributed by client IO. See the LICENSE file.*/


import {
    CONFIRMATION_ICON,
    ENTITY_ICON,
    MESSAGE_ICON,
    USER_INPUT_ICON
} from '../../theme';

export const stencilConfig = {
    shapes: [{
        name: 'FlowchartStart'
    }, {
        name: 'FlowchartEnd'
    }, {
        name: 'Message',
        attrs: {
            label: { text: 'API Request' },
            icon: { xlinkHref: USER_INPUT_ICON }
        }
    }, {
        name: 'Message2',
        attrs: {
            label: { text: 'LLM Format' },
            icon: { xlinkHref: USER_INPUT_ICON }
        }
    }, {
        name: 'Message3',
        attrs: {
            label: { text: 'LLM Translate' },
            icon: { xlinkHref: USER_INPUT_ICON }
        }
    }, {
        name: 'Message4',
        attrs: {
            label: { text: 'LLM Generate' },
            icon: { xlinkHref: USER_INPUT_ICON }
        }
    }, {
        name: 'Message5',
        attrs: {
            label: { text: 'LLM Keyword Generate' },
            icon: { xlinkHref: USER_INPUT_ICON }
        }
    }, {
        name: 'Message6',
        attrs: {
            label: { text: 'LLM Image Prompt Generate' },
            icon: { xlinkHref: USER_INPUT_ICON }
        }
    }, {
        name: 'Message7',
        attrs: {
            label: { text: 'LLM Image Generate' },
            icon: { xlinkHref: USER_INPUT_ICON }
        }
    }, {
        name: 'Message8',
        attrs: {
            label: { text: 'Get Document' },
            icon: { xlinkHref: USER_INPUT_ICON }
        }
    }, {
        name: 'Message9',
        attrs: {
            label: { text: 'Display Block Output' },
            icon: { xlinkHref: USER_INPUT_ICON }
        }
    }, {
        name: 'Constant',
        attrs: {
            label: { text: 'Constant Value' },
            icon: { xlinkHref: USER_INPUT_ICON }
        }
    }]
};

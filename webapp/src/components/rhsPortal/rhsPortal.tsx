// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React, {useState, useLayoutEffect} from 'react'
import ReactDOM from 'react-dom'

type Props = {
    children: React.ReactNode
}

import './rhsPortal.scss'

const RhsPortal = (props: Props): JSX.Element => {
    const [el] = useState(document.createElement('div', ))
    el.classList.add('RHSPortal')

    const rootPortal = document.getElementById('focalboard-rhs-portal')

    useLayoutEffect(() => {
        if (rootPortal) {
            rootPortal.appendChild(el)
        }
        return () => {
            if (rootPortal) {
                rootPortal.removeChild(el)
            }
        }
    }, [])

    return ReactDOM.createPortal(props.children, el)  // eslint-disable-line
}

export default React.memo(RhsPortal)

import React from 'react'
import Form from './index'

const NewEdition = ({form, onChange, onSubmit}) => (
    <div className="NewEdition_Lateral_Spaces row">

        <div className="col-sm NewEdition_Form_Space">
            <Form
                onChange={onChange}
                onSubmit={onSubmit}
                form={form}
            />            
        </div>
    </div>
)

export default NewEdition
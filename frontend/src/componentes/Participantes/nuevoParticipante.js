import React from 'react'
import Form from './index'

const NewParticipante = ({form, onChange, onSubmit}) => (
    <div className="NewParticipante_Lateral_Spaces row">

        <div className="col-sm NewParticipante_Form_Space">
            <Form
                onChange={onChange}
                onSubmit={onSubmit}
                form={form}
            />            
        </div>
    </div>
)

export default NewParticipante
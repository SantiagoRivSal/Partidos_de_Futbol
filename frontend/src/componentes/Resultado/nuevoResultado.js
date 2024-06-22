import React from 'react';
import Form from './index';

const NewResultado = ({form, onChange, onSubmit, equipos}) => (
  <div className="NewParticipante_Lateral_Spaces row">
    <div className="col-sm NewParticipante_Form_Space">
      <Form
        onChange={onChange}
        onSubmit={onSubmit}
        form={form}
        equipos={equipos}
      />            
    </div>
  </div>
);

export default NewResultado;

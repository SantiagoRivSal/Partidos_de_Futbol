import React from "react";
import { Routes, Route } from "react-router-dom";
import { PaisesLista } from "./Equipos/index";
import { InsertEdition } from "./Torneos/funcion";
import { InsertParticipante } from "./Participantes/funcion";
import { MenuOpciones } from "./Ediciones/Opciones";
import { TorneosLista } from "./Ediciones/index";
import { Match } from "./Partidos/index";
import { InsertResultado } from "./Resultado/funcion";


export const Paginas = () => {
  return (
    <section>
        <Routes>
          <Route path="/" element={<PaisesLista />} />
          <Route path="/torneos" element={<InsertEdition />} />
          <Route path="/participantes" element={<InsertParticipante />} />
          <Route path="/ediciones/opciones" element={<MenuOpciones />} />
          <Route path="/ediciones" element={<TorneosLista />} />
          <Route path="/participantes/edicion" element={<Match />} />
          <Route path="/resultado" element={<InsertResultado />} />
        </Routes>
    </section>
)
};

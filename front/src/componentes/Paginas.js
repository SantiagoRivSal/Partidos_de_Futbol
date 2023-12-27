import React from "react";
import { Routes, Route } from "react-router-dom";
import { PaisesLista } from "./Equipos/index";
import { InsertEdition } from "./Torneos/funcion";
import { Home } from "./Home/index";
import { InsertParticipante } from "./Participantes/funcion";
import { TorneosLista } from "./Ediciones/index";

export const Paginas = () => {
  return (
    <section>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/equipos" element={<PaisesLista />} />
          <Route path="/torneos" element={<InsertEdition />} />
          <Route path="/participantes" element={<InsertParticipante />} />
          <Route path="/ediciones" element={<TorneosLista />} />
        </Routes>
    </section>
)
};

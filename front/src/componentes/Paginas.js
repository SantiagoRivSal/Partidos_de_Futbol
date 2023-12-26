import React from "react";
import { Routes, Route } from "react-router-dom";
import { PaisesLista } from "./Equipos/index";
import { InsertEdition } from "./Torneos/funcion";
import { Home } from "./Home/index";
import { Historiales } from "./Historiales/index";
import { TorneosLista } from "./Partidos/index";

export const Paginas = () => {
  return (
    <section>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/equipos" element={<PaisesLista />} />
          <Route path="/torneos" element={<InsertEdition />} />
          <Route path="/historiales" element={<Historiales />} />
          <Route path="/ediciones" element={<TorneosLista />} />
        </Routes>
    </section>
)
};

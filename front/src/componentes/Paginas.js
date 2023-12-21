import React from "react"
import {Routes, Route} from "react-router-dom";
import {Equipos} from "./Equipos/index";
import {Torneos} from "./Torneos/index";
import {Home} from "./Home/index"
import {Historiales} from "./Historiales/index";

export const Paginas = ()=>{
    return(
        <section>
          <Routes>
          <Route path="/" exact element={<Home/>} />
        <Route path="/equipos" exact element={<Equipos/>} />
        <Route path="/torneos" exact element={<Torneos/>} />
        <Route path="/historiales"  element={<Historiales/>} />
        </Routes>  
        </section>
    )
}

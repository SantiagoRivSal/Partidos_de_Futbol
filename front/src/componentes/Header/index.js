import React from "react"
import Logo from "../../imagenes/fifa.png"
import './header.css';
import {Link} from "react-router-dom";
  

export const Header = ()=>{
    
    return(
        <header>
        <Link to="/">
            <div className="logo">
                <img src={Logo} alt="logo" width="150"/>
            </div>
        </Link>
        
        <ul>
            <li>
               <Link to="/" className="botones_menu">
               INICIO
               </Link> 
            </li>
            <li>
               <Link to="/equipos" className="botones_menu">
               EQUIPOS
               </Link> 
            </li>
            <li>
               <Link to="/historiales" className="botones_menu"> 
                HISTORIALES
                </Link> 
            </li>
            <li>
               <Link to="/ediciones" className="botones_menu"> 
                EDICIONES
                </Link> 
            </li>
            <li>
               <Link to="/torneos" className="botones_menu"> 
                TORNEOS
                </Link> 
            </li>
        </ul>
        </header>
    )
}
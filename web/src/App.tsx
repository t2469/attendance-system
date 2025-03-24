import React from 'react'
import './App.css'
import EmployeeList from './pages/EmployeeList.tsx'
import EmployeeDetail from './pages/EmployeeDetail.tsx'
import Login from './pages/Login.tsx'
import {BrowserRouter, Route, Routes} from "react-router-dom";

const App: React.FC = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<EmployeeList/>}/>
                <Route path="/login" element={<Login/>}/>
                <Route path="/employees/:id" element={<EmployeeDetail/>}/>
            </Routes>
        </BrowserRouter>
    );
};

export default App

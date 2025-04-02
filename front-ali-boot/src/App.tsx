import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { Navbar } from './components/Navbar';
import { StudentWelcome } from './pages/StudentWelcome';
import { CollegeWelcome } from './pages/CollegeWelcome';
import { Register } from './pages/Register';

function App() {
  return (
    <BrowserRouter>
      <div className="min-vh-100 bg-light">
        <Navbar />
        <Routes>
          <Route path="/" element={<StudentWelcome />} />
          <Route path="/college" element={<CollegeWelcome />} />
          <Route path="/search" element={<div className="container py-5"><h1>College Search</h1></div>} />
          <Route path="/register" element={<Register />} />
        </Routes>
      </div>
    </BrowserRouter>
  );
}

export default App;
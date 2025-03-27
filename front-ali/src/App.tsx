import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { Navbar } from './components/Navbar';
import { StudentWelcome } from './pages/StudentWelcome';
import { ProfessionalWelcome } from './pages/ProfessionalWelcome';

function App() {
  return (
    <BrowserRouter>
      <div className="min-h-screen bg-gray-50">
        <Navbar />
        <Routes>
          <Route path="/" element={<StudentWelcome />} />
          <Route path="/professional" element={<ProfessionalWelcome />} />
        </Routes>
      </div>
    </BrowserRouter>
  );
}

export default App;
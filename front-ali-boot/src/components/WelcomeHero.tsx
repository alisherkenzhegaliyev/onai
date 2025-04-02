import React from 'react';
import { Search } from 'lucide-react';
import { Link } from 'react-router-dom';

export const WelcomeHero: React.FC = () => (
  <div className="position-relative overflow-hidden">
    <div className="position-absolute w-100 h-100">
      <img
        src="https://images.unsplash.com/photo-1541339907198-e08756dedf3f?auto=format&fit=crop&q=80&w=2070"
        alt="College campus"
        className="w-100 h-100 object-fit-cover opacity-25"
      />
    </div>
    <div className="position-relative py-5">
      <div className="container text-center py-5">
        <p className="text-primary text-uppercase fw-semibold mb-2">
          One Application - Infinite Opportunities
        </p>
        <h1 className="display-3 fw-bold mb-4">
          Your Gateway to Higher Education
        </h1>
        <p className="lead mb-5 mx-auto" style={{ maxWidth: '600px' }}>
          Submit one application to multiple colleges. Save time, reduce stress, 
          and maximize your chances of acceptance with our streamlined process.
        </p>
        <div className="d-flex justify-content-center gap-3">
          <button className="btn btn-primary btn-lg px-4">
            Start Application
          </button>
          <Link to="/search" className="btn btn-outline-dark btn-lg px-4 d-flex align-items-center">
            <Search className="me-2" style={{ width: '20px', height: '20px' }} />
            Find Colleges
          </Link>
        </div>
      </div>
    </div>
  </div>
);
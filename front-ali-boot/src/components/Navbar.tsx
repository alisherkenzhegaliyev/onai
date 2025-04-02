import React from 'react';
import { Link, useLocation } from 'react-router-dom';
import { Logo } from './Logo';
import { AuthButtons } from './AuthButtons';
import { Building2 } from 'lucide-react';

export const Navbar: React.FC = () => {
  const location = useLocation();
  const isCollege = location.pathname === '/college';

  return (
    <nav className="navbar navbar-expand-lg navbar-light bg-white shadow-sm">
      <div className="container">
        <Link to="/" className="navbar-brand">
          <Logo />
        </Link>
        
        <div className="navbar-collapse" id="navbarContent">
          <div className="navbar-nav me-auto mb-2 mb-lg-0">
            {/* Navigation links removed as requested */}
          </div>
          
          <div className="d-flex align-items-center gap-3">
            <Link 
              to={isCollege ? '/' : '/college'} 
              className="nav-link d-flex align-items-center"
            >
              <Building2 className="me-2" style={{ width: '16px', height: '16px' }} />
              {isCollege ? 'Student Portal' : 'For Colleges'}
            </Link>
            <AuthButtons />
          </div>
        </div>
      </div>
    </nav>
  );
};
import React from 'react';
import { Link, useLocation } from 'react-router-dom';
import { Logo } from './Logo';
import { NavLink } from './NavLink';
import { AuthButtons } from './AuthButtons';
import { Building2 } from 'lucide-react';

export const Navbar: React.FC = () => {
  const location = useLocation();
  const isCollege = location.pathname === '/college';

  return (
    <nav className="bg-white shadow-md">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between items-center h-16">
          <div className="flex items-center">
            <Logo />
            <div className="hidden md:block">
              <div className="ml-10 flex items-baseline space-x-4">
                <NavLink href="#">{isCollege ? 'BROWSE APPLICATIONS' : 'FIND COLLEGES'}</NavLink>
              </div>
            </div>
          </div>
          <div className="flex items-center space-x-4">
            <Link 
              to={isCollege ? '/' : '/college'} 
              className="flex items-center text-gray-700 hover:text-gray-900 px-3 py-2 rounded-md text-sm font-medium"
            >
              <Building2 className="h-4 w-4 mr-2" />
              {isCollege ? 'Student Portal' : 'For Colleges'}
            </Link>
            <AuthButtons />
          </div>
        </div>
      </div>
    </nav>
  );
};
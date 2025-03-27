import React from 'react';
import { Link } from 'react-router-dom';

interface NavLinkProps {
  href: string;
  children: React.ReactNode;
}

export const NavLink: React.FC<NavLinkProps> = ({ href, children }) => (
  <Link
    to={href}
    className="text-gray-900 hover:bg-gray-100 px-3 py-2 rounded-md text-sm font-medium"
  >
    {children}
  </Link>
);
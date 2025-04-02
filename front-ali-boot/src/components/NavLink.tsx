import React from 'react';
import { Link } from 'react-router-dom';

interface NavLinkProps {
  href: string;
  children: React.ReactNode;
}

export const NavLink: React.FC<NavLinkProps> = ({ href, children }) => (
  <Link
    to={href}
    className="nav-link px-3 py-2 d-flex align-items-center"
  >
    {children}
  </Link>
);
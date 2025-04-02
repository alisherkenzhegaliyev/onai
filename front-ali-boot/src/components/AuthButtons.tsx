import React, { useState } from 'react';
import { LogIn, UserPlus } from 'lucide-react';
import { Link } from 'react-router-dom';
import { SignInModal } from './SignInModal';

export const AuthButtons: React.FC = () => {
  const [isSignInOpen, setIsSignInOpen] = useState(false);

  return (
    <div className="d-flex align-items-center gap-3">
      <button 
        onClick={() => setIsSignInOpen(true)}
        className="btn btn-link text-decoration-none d-flex align-items-center"
      >
        <LogIn className="me-2" style={{ width: '16px', height: '16px' }} />
        Sign in
      </button>
      <Link to="/register" className="btn btn-dark d-flex align-items-center">
        <UserPlus className="me-2" style={{ width: '16px', height: '16px' }} />
        Create an account
      </Link>
      <SignInModal 
        isOpen={isSignInOpen}
        onClose={() => setIsSignInOpen(false)}
      />
    </div>
  );
};
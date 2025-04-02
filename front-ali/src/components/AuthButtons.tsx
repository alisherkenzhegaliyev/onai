import React, { useState } from 'react';
import { LogIn, UserPlus } from 'lucide-react';
import { SignInModal } from './SignInModal';

export const AuthButtons: React.FC = () => {
  const [isSignInOpen, setIsSignInOpen] = useState(false);

  return (
    <div className="flex items-center space-x-4">
      <button 
        onClick={() => setIsSignInOpen(true)}
        className="flex items-center text-gray-700 hover:text-gray-900 px-3 py-2 rounded-md text-sm font-medium"
      >
        <LogIn className="h-4 w-4 mr-2" />
        Sign in
      </button>
      <button className="flex items-center bg-gray-900 text-white hover:bg-gray-800 px-4 py-2 rounded-md text-sm font-medium">
        <UserPlus className="h-4 w-4 mr-2" />
        Create an account
      </button>
      <SignInModal 
        isOpen={isSignInOpen}
        onClose={() => setIsSignInOpen(false)}
      />
    </div>
  );
};
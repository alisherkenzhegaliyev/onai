import React from 'react';
import { Building2 } from 'lucide-react';

export const ProfessionalWelcome: React.FC = () => (
  <div className="hero-container">
    <div className="hero-background">
      <img
        src="https://images.unsplash.com/photo-1486406146926-c627a92ad1ab?auto=format&fit=crop&q=80&w=2070"
        alt="Professional building"
      />
    </div>
    <div className="hero-content">
      <div className="hero-text-container">
        <h2 className="hero-tagline">
          One Platform - Endless Talent
        </h2>
        <h1 className="mt-4">
          Your Gateway to Top Graduate Talent
        </h1>
        <p className="hero-description">
          Connect with qualified candidates from top institutions. Streamline your recruitment process 
          and find the perfect match for your organization.
        </p>
        <div className="hero-buttons">
          <button className="button-primary">
            Post a Position
          </button>
          <button className="button-secondary">
            Browse Candidates <span aria-hidden="true">→</span>
          </button>
        </div>
      </div>
    </div>
  </div>
);
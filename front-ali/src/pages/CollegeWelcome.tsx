import React from 'react';
import { Building2 } from 'lucide-react';

export const CollegeWelcome: React.FC = () => (
  <div className="hero-container">
    <div className="hero-background">
      <img
        src="https://images.unsplash.com/photo-1607237138185-eedd9c632b0b?auto=format&fit=crop&q=80&w=2070"
        alt="College campus buildings"
      />
    </div>
    <div className="hero-content">
      <div className="hero-text-container">
        <h2 className="hero-tagline">
          One Platform - Exceptional Students
        </h2>
        <h1 className="mt-4">
          Your Gateway to Tomorrow's Leaders
        </h1>
        <p className="hero-description">
          Join our network of prestigious institutions. Access a curated pool of motivated students 
          and streamline your admissions process through our comprehensive platform.
        </p>
        <div className="hero-buttons">
          <button className="button-primary">
            Become a Partner College
          </button>
          <button className="button-secondary">
            Browse Applications <span aria-hidden="true">→</span>
          </button>
        </div>
      </div>
    </div>
  </div>
);
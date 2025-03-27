import React from 'react';

export const WelcomeHero: React.FC = () => (
  <div className="hero-container">
    <div className="hero-background">
      <img
        src="https://images.unsplash.com/photo-1541339907198-e08756dedf3f?auto=format&fit=crop&q=80&w=2070"
        alt="College campus"
      />
    </div>
    <div className="hero-content">
      <div className="hero-text-container">
        <h2 className="hero-tagline">
          One Application - Infinite Opportunities
        </h2>
        <h1 className="mt-4">
          Your Gateway to Higher Education
        </h1>
        <p className="hero-description">
          Submit one application to multiple colleges. Save time, reduce stress, 
          and maximize your chances of acceptance with our streamlined process.
        </p>
        <div className="hero-buttons">
          <button className="button-primary">
            Start Application
          </button>
          <button className="button-secondary">
            Start Your Search <span aria-hidden="true">→</span>
          </button>
        </div>
      </div>
    </div>
  </div>
);
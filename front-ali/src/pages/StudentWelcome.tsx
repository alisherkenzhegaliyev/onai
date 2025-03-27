import React from 'react';
import { WelcomeHero } from '../components/WelcomeHero';
import { FeatureSection } from '../components/FeatureSection';

export const StudentWelcome: React.FC = () => (
  <main>
    <WelcomeHero />
    <FeatureSection />
  </main>
);
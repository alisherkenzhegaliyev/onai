import React from 'react';
import { Search, GraduationCap, BookOpen } from 'lucide-react';

const features = [
  {
    name: 'Smart Search',
    description: 'Find colleges that match your preferences, from location to program specifics.',
    icon: Search,
  },
  {
    name: 'Compare Programs',
    description: 'Side-by-side comparison of programs, tuition costs, and admission requirements.',
    icon: BookOpen,
  },
  {
    name: 'Application Tracking',
    description: 'Keep track of your college applications, deadlines, and requirements in one place.',
    icon: GraduationCap,
  },
];

export const FeatureSection: React.FC = () => (
  <div className="py-5">
    <div className="container py-5">
      <div className="text-center mb-5">
        <h6 className="text-primary text-uppercase fw-semibold">Start Your Journey</h6>
        <h2 className="display-5 fw-bold">
          Everything you need to find your dream college
        </h2>
      </div>
      
      <div className="row g-4">
        {features.map((feature) => (
          <div key={feature.name} className="col-md-4">
            <div className="h-100 p-4">
              <div className="d-flex align-items-center mb-3">
                <feature.icon className="text-primary me-2" style={{ width: '24px', height: '24px' }} />
                <h5 className="mb-0">{feature.name}</h5>
              </div>
              <p className="text-muted mb-0">{feature.description}</p>
            </div>
          </div>
        ))}
      </div>
    </div>
  </div>
);
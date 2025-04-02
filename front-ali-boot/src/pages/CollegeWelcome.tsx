import React from 'react';
import { GraduationCap, Users, BarChart3, Clock } from 'lucide-react';

const features = [
  {
    name: 'Streamlined Applications',
    description: 'Access and review student applications through our intuitive dashboard.',
    icon: GraduationCap,
  },
  {
    name: 'Student Analytics',
    description: 'Gain valuable insights into your applicant pool with detailed analytics.',
    icon: BarChart3,
  },
  {
    name: 'Efficient Processing',
    description: 'Reduce administrative workload with our automated processing system.',
    icon: Clock,
  },
];

export const CollegeWelcome: React.FC = () => (
  <main>
    <div className="position-relative overflow-hidden">
      <div className="position-absolute w-100 h-100">
        <img
          src="https://images.unsplash.com/photo-1607237138185-eedd9c632b0b?auto=format&fit=crop&q=80&w=2070"
          alt="College campus buildings"
          className="w-100 h-100 object-fit-cover opacity-25"
        />
      </div>
      <div className="position-relative py-5">
        <div className="container text-center py-5">
          <p className="text-primary text-uppercase fw-semibold mb-2">
            One Platform - Exceptional Students
          </p>
          <h1 className="display-3 fw-bold mb-4">
            Your Gateway to Tomorrow's Leaders
          </h1>
          <p className="lead mb-5 mx-auto" style={{ maxWidth: '600px' }}>
            Join our network of prestigious institutions. Access a curated pool of motivated students 
            and streamline your admissions process through our comprehensive platform.
          </p>
          <div className="d-flex justify-content-center gap-3">
            <button className="btn btn-primary btn-lg px-4">
              Become a Partner College
            </button>
            <button className="btn btn-outline-dark btn-lg px-4 d-flex align-items-center">
              <GraduationCap className="me-2" style={{ width: '20px', height: '20px' }} />
              Browse Applications
            </button>
          </div>
        </div>
      </div>
    </div>

    <div className="py-5 bg-white">
      <div className="container py-5">
        <div className="text-center mb-5">
          <h6 className="text-primary text-uppercase fw-semibold">Why Partner With Us</h6>
          <h2 className="display-5 fw-bold">
            Streamline Your Admissions Process
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
  </main>
);
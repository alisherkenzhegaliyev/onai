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
  <div className="py-16 sm:py-24">
    <div className="mx-auto max-w-7xl px-6 lg:px-8">
      <div className="mx-auto max-w-2xl lg:text-center">
        <h2 className="text-base font-semibold leading-7 text-sky-600">Start Your Journey</h2>
        <p className="mt-2 text-3xl font-bold tracking-tight text-gray-900 sm:text-4xl">
          Everything you need to find your dream college
        </p>
      </div>
      <div className="mx-auto mt-16 max-w-2xl sm:mt-20 lg:mt-24 lg:max-w-none">
        <dl className="grid max-w-xl grid-cols-1 gap-x-8 gap-y-16 lg:max-w-none lg:grid-cols-3">
          {features.map((feature) => (
            <div key={feature.name} className="flex flex-col">
              <dt className="flex items-center gap-x-3 text-base font-semibold leading-7 text-gray-900">
                <feature.icon className="h-5 w-5 flex-none text-sky-600" aria-hidden="true" />
                {feature.name}
              </dt>
              <dd className="mt-4 flex flex-auto flex-col text-base leading-7 text-gray-600">
                <p className="flex-auto">{feature.description}</p>
              </dd>
            </div>
          ))}
        </dl>
      </div>
    </div>
  </div>
);
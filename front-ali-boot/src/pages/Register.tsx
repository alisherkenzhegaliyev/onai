import React, { useState } from 'react';

export const Register: React.FC = () => {
  const [formData, setFormData] = useState({
    email: '',
    emailConfirm: '',
    password: '',
    passwordConfirm: '',
    firstName: '',
    lastName: '',
    dateOfBirth: '',
    country: '',
    acceptTerms: false
  });

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // Handle form submission
  };

  return (
    <div className="container py-5">
      <div className="row justify-content-center">
        <div className="col-md-8 col-lg-6">
          <h2 className="text-center mb-4">Create Your Account</h2>
          <div className="card shadow-sm">
            <div className="card-body p-4">
              <form onSubmit={handleSubmit}>
                <div className="mb-4">
                  <h5 className="card-title mb-3">Account Information</h5>
                  <div className="mb-3">
                    <label htmlFor="email" className="form-label">Email address*</label>
                    <input
                      type="email"
                      className="form-control"
                      id="email"
                      value={formData.email}
                      onChange={(e) => setFormData({...formData, email: e.target.value})}
                      required
                    />
                  </div>
                  <div className="mb-3">
                    <label htmlFor="emailConfirm" className="form-label">Confirm email address*</label>
                    <input
                      type="email"
                      className="form-control"
                      id="emailConfirm"
                      value={formData.emailConfirm}
                      onChange={(e) => setFormData({...formData, emailConfirm: e.target.value})}
                      required
                    />
                  </div>
                  <div className="mb-3">
                    <label htmlFor="password" className="form-label">Password*</label>
                    <input
                      type="password"
                      className="form-control"
                      id="password"
                      value={formData.password}
                      onChange={(e) => setFormData({...formData, password: e.target.value})}
                      required
                      minLength={8}
                    />
                    <div className="form-text">
                      Must be at least 8 characters long
                    </div>
                  </div>
                  <div className="mb-3">
                    <label htmlFor="passwordConfirm" className="form-label">Confirm password*</label>
                    <input
                      type="password"
                      className="form-control"
                      id="passwordConfirm"
                      value={formData.passwordConfirm}
                      onChange={(e) => setFormData({...formData, passwordConfirm: e.target.value})}
                      required
                      minLength={8}
                    />
                  </div>
                </div>

                <div className="mb-4">
                  <h5 className="card-title mb-3">Personal Information</h5>
                  <div className="mb-3">
                    <label htmlFor="firstName" className="form-label">Legal first name*</label>
                    <input
                      type="text"
                      className="form-control"
                      id="firstName"
                      value={formData.firstName}
                      onChange={(e) => setFormData({...formData, firstName: e.target.value})}
                      required
                    />
                  </div>
                  <div className="mb-3">
                    <label htmlFor="lastName" className="form-label">Legal last name*</label>
                    <input
                      type="text"
                      className="form-control"
                      id="lastName"
                      value={formData.lastName}
                      onChange={(e) => setFormData({...formData, lastName: e.target.value})}
                      required
                    />
                  </div>
                  <div className="mb-3">
                    <label htmlFor="dateOfBirth" className="form-label">Date of birth*</label>
                    <input
                      type="date"
                      className="form-control"
                      id="dateOfBirth"
                      value={formData.dateOfBirth}
                      onChange={(e) => setFormData({...formData, dateOfBirth: e.target.value})}
                      required
                    />
                  </div>
                  <div className="mb-3">
                    <label htmlFor="country" className="form-label">Country of residence*</label>
                    <select
                      className="form-select"
                      id="country"
                      value={formData.country}
                      onChange={(e) => setFormData({...formData, country: e.target.value})}
                      required
                    >
                      <option value="">Select a country</option>
                      <option value="US">United States</option>
                      <option value="CA">Canada</option>
                      <option value="GB">United Kingdom</option>
                      <option value="AU">Australia</option>
                      <option value="NZ">New Zealand</option>
                      {/* Add more countries as needed */}
                    </select>
                  </div>
                </div>

                <div className="mb-4">
                  <div className="form-check">
                    <input
                      type="checkbox"
                      className="form-check-input"
                      id="acceptTerms"
                      checked={formData.acceptTerms}
                      onChange={(e) => setFormData({...formData, acceptTerms: e.target.checked})}
                      required
                    />
                    <label className="form-check-label" htmlFor="acceptTerms">
                      I agree to the processing of my personal information according to the Privacy Policy*
                    </label>
                  </div>
                </div>

                <button type="submit" className="btn btn-primary w-100 btn-lg">
                  Create Account
                </button>

                <div className="mt-3 text-center text-muted">
                  <small>* Required fields</small>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};
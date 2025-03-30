// PackagesList.test.js
import React from 'react';
import { render, screen, waitFor } from '@testing-library/react';
import { BrowserRouter } from 'react-router-dom';
import PackagesList from './PackagesList';

jest.mock('../services/packageService', () => ({
  getPackages: jest.fn(),
}));

import { getPackages } from '../services/packageService';

const mockPackages = [
  {
    id: 1,
    package_name: "Explore New York",
    package_description: "Enjoy the vibrant city life.",
    package_price: 999,
    days: 5,
    nights: 4,
    location: "New York",
  }
];

test('renders available packages correctly', async () => {
  getPackages.mockResolvedValue(mockPackages);

  render(
    <BrowserRouter>
      <PackagesList />
    </BrowserRouter>
  );

  expect(screen.getByText(/Available Packages/i)).toBeInTheDocument();

  await waitFor(() => {
    expect(screen.getByText("Explore New York")).toBeInTheDocument();
    expect(screen.getByText("Enjoy the vibrant city life.")).toBeInTheDocument();
    expect(screen.getByText("999")).toBeInTheDocument();
  });
});

import { createRoot } from 'react-dom/client';
import { router } from './Router'
import { RouterProvider } from 'react-router-dom';

const root = createRoot(document.getElementById('root')!);
root.render(
  <RouterProvider router={router} />
)

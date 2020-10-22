import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import CreatePatientRecord from './components/Users';
 
export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/user', CreatePatientRecord);
  },
});


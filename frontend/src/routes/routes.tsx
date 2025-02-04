import { Routes, Route } from 'react-router-dom';
import Login from '../pages/login/Login';
import Register from '../pages/register/Register';
import Healthcheck from '@/pages/healthcheck/Healthcheck';

const AppRoutes = () => {
  return (
    <Routes>
      <Route path='/' element={<Healthcheck />} />
      <Route path='login' element={<Login />} />
      <Route path='register' element={<Register />} />
    </Routes>
  )
};

export default AppRoutes;

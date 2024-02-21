import { Route, Routes} from "react-router-dom";
import {
  Home,
  Payment,
  Employee,
  User,
  Ticket
} from "./page";

function App() {
  return (
    <Routes>
      {/* <Route path="" element={<Home/>} /> */}
      {/* <Route path="/payment" element={<Payment/>} /> */}
      <Route path="/employee" element={<Employee/>} />
      {/* <Route path="/user" element={<User/>} />
      <Route path="/ticket" element={<Ticket/>} /> */}
    </Routes>
  );
};

export default App;

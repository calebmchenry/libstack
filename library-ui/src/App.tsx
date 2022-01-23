import React from "react";
import { Login } from "./pages/Login";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { Home } from "./pages/Home";
import { Layout } from "./Layout";
import { auth } from "./auth";
import { SignUp } from "./pages/SignUp";

function App() {
  return (
    <auth.Provider>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Layout />}>
            <Route index element={<Home />} />
            <Route path="login" element={<Login />} />
            <Route path="sign-up" element={<SignUp />} />
          </Route>
        </Routes>
      </BrowserRouter>
    </auth.Provider>
  );
}

export default App;

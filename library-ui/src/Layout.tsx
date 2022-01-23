import React from "react";
import { Outlet } from "react-router-dom";

export function Layout() {
  return (
    <div>
      <header>
        <span>Libstack📚</span>
      </header>
      <main>
        <Outlet />
      </main>
    </div>
  );
}

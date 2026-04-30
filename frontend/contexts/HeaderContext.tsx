"use client";

import React, { createContext, useContext, useState, ReactNode } from "react";

interface HeaderContextProps {
  title: string;
  isMenuOpen: boolean;
  setTitle: (newTitle: string) => void;
  toggleMenu: () => void;
}

const HeaderContext = createContext<HeaderContextProps | undefined>(undefined);

export function HeaderProvider({ children }: { children: ReactNode }) {
  const [title, setTitle] = useState("Default Header Title");
  const [isMenuOpen, setIsMenuOpen] = useState(false);

  const toggleMenu = () => setIsMenuOpen((prev) => !prev);

  return (
    <HeaderContext.Provider value={{ title, isMenuOpen, setTitle, toggleMenu }}>
      {children}
    </HeaderContext.Provider>
  );
}

// Custom hook để dùng context
export function useHeader() {
  const context = useContext(HeaderContext);
  if (!context) {
    throw new Error("useHeader must be used within a HeaderProvider");
  }
  return context;
}

"use client";
import { useHeader } from "@/contexts/HeaderContext";
import { SidebarTrigger } from "./ui/sidebar";

export default function Header() {
  const { title, isMenuOpen, toggleMenu } = useHeader();

  return (
    <header className="flex items-center justify-between p-4 bg-gray-100">
      <div className="flex flex-row items-center justify-items-start gap-4">
        <SidebarTrigger className="-ml-1" />
        <h1 className="text-xl font-bold">{title}</h1>
      </div>
      <button onClick={toggleMenu}>
        {isMenuOpen ? "Close Menu" : "Open Menu"}
      </button>
    </header>
  );
}
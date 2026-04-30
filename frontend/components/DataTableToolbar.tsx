"use client";

import { Input } from "@/components/ui/input";

interface DataTableToolbarProps {
  search: string;
  onSearchChange: (v: string) => void;
}

export default function DataTableToolbar({ search, onSearchChange }: DataTableToolbarProps) {
  return (
    <div className="py-2">
      <Input
        placeholder="Search..."
        value={search}
        onChange={(e) => onSearchChange(e.target.value)}
        className="w-72"
      />
    </div>
  );
}

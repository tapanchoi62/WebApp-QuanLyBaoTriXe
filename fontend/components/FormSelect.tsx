"use client";

import * as React from "react";
import { Control, Controller } from "react-hook-form";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Label } from "./ui/label";

interface Option {
  label: string;
  value: string;
}

interface FormSelectProps {
  name: string;
  control: Control<any>;
  options: Option[];
  placeholder?: string;
  isRequired?: boolean;
  direction?: "vertical" | "horizontal";
  label?: string;
}

export function FormSelect({
  name,
  control,
  options,
  label,
  placeholder = "Select an option",
  isRequired = false,
  direction = "vertical",
}: FormSelectProps) {
  return (
    <Controller
      name={name}
      control={control}
      rules={isRequired ? { required: `${placeholder} is required` } : {}}
      render={({ field, fieldState }) => (
        <div className={direction === "vertical" ? "flex flex-col gap-2 w-full" : "flex items-center gap-2 w-full"}>
          <Label className="w-full">
            {label} {isRequired && <span className="text-red-500">*</span>}
          </Label>
          <Select  value={field.value} onValueChange={field.onChange}>
            <SelectTrigger className="w-full">
              <SelectValue placeholder={placeholder} />
            </SelectTrigger>
            <SelectContent>
              {options.map((opt) => (
                <SelectItem key={opt.value} value={opt.value}>
                  {opt.label}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
          {fieldState.error && (
            <p className="text-red-500 text-sm mt-1">
              {fieldState.error.message}
            </p>
          )}
        </div>
      )}
    />
  );
}

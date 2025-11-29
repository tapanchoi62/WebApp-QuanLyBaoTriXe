"use client";

import  { ReactNode } from "react";
import { FormProvider, UseFormReturn } from "react-hook-form";

interface FormHookProviderProps<T> {
  children: ReactNode;
  form: UseFormReturn<any>;
  onSubmit: any;
}

export function FormHookProvider<T>({
  children,
  onSubmit,
  form
}: FormHookProviderProps<T>) {

  return (
    <FormProvider {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4 p-6">
        {children}
      </form>
    </FormProvider>
  );
}

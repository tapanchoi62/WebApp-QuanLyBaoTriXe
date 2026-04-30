'use client';

import { useFormContext, Controller } from 'react-hook-form';
import { Label } from '@/components/ui/label';
import { Input } from '@/components/ui/input';
import { cn } from '@/lib/utils'; // helper classnames, optional

interface FormInputProps {
  name: string;
  label?: string;
  placeholder?: string;
  type?: string;
  disabled?: boolean;
  className?: string;
  readOnly?: boolean;
}

export function FormInput({
  name,
  label,
  placeholder,
  type = 'text',
  className,
  disabled = false,
  readOnly = false,
}: FormInputProps) {
  const { control } = useFormContext();

  return (
    <div className={cn('flex flex-col space-y-1', className)}>
      {label && <Label htmlFor={name}>{label}</Label>}
      <Controller
        name={name}
        control={control}
        render={({ field, fieldState }) => (
          <>
            <Input
              readOnly={readOnly}
              id={name}
              type={type}
              placeholder={placeholder}
              disabled={disabled}
              {...field}
            />
            {fieldState.error && (
              <span className="text-red-500 text-sm">{fieldState.error.message}</span>
            )}
          </>
        )}
      />
    </div>
  );
}

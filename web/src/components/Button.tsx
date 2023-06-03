import React from "react";
import clsx from "clsx";

export type ButtonProps = JSX.IntrinsicElements["button"] & {};

export const Button: React.FC<ButtonProps> = ({
  className,
  children,
  ...props
}) => {
  return (
    <button
      className={clsx(
        "border border-primary rounded py-1 px-2 transition duration-200 hover:bg-accent hover:text-white",
        className
      )}
      {...props}
    >
      {children}
    </button>
  );
};

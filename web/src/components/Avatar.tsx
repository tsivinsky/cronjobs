import clsx from "clsx";
import Image, { ImageProps } from "next/image";
import React from "react";

export type AvatarProps = Omit<
  ImageProps,
  "width" | "height" | "layout" | "fill"
> & {
  size?: number;
};

export const Avatar: React.FC<AvatarProps> = ({
  src,
  size = 32,
  className,
  ...props
}) => {
  return (
    <div
      className={clsx("relative rounded-full overflow-hidden", className)}
      style={{ width: `${size}px`, height: `${size}px` }}
    >
      <Image src={src} fill className="scale-150" {...props} />
    </div>
  );
};

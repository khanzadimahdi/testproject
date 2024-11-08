"use client";
import Image from "next/image";
import {useState} from "react";
import {Avatar, Skeleton} from "@mantine/core";
import {useInit} from "@/hooks/data/init";
import {FILES_PUBLIC_URL} from "@/constants/envs";
import classes from "./auth-user-avatar.module.css";

type Props = {
  width?: number;
  height?: number;
};

export function AuthUserAvatar({width = 45, height = 45}: Props) {
  const {data, isLoading} = useInit();
  const [hasImageFailed, setHasImageFailed] = useState(false);
  if (isLoading) {
    return <Skeleton circle width={width} height={height} />;
  }

  if (hasImageFailed) {
    return <Avatar src={null} w={width} h={height} />;
  }

  if (data?.status === "authenticated") {
    const {avatar, name} = data.profile;
    return (
      <Image
        src={`${FILES_PUBLIC_URL}/${avatar}`}
        alt={name}
        width={width}
        height={height}
        className={classes.avatar}
        onError={() => setHasImageFailed(true)}
      />
    );
  }
  return null;
}
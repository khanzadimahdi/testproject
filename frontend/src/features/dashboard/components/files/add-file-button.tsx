"use client";
import {useState, useRef, useId} from "react";
import {Tooltip, ActionIcon, Loader} from "@mantine/core";
import {IconPlus} from "@tabler/icons-react";
import {addFileAction} from "../../actions/add-file";

export function AddFileButton() {
  const [pending, setPending] = useState(false);
  const formRef = useRef<HTMLFormElement>(null);
  const inputRef = useRef<HTMLInputElement>(null);
  const inputId = useId();

  const handleFileChange = async () => {
    if (formRef.current) {
      const fd = new FormData(formRef.current);
      try {
        setPending(true);
        await addFileAction(fd);
      } finally {
        setPending(false);
      }
    }
  };

  return (
    <form ref={formRef}>
      <Tooltip
        label={pending ? "در حال اضافه کردن فایل ها" : "اضافه کردن تصویر"}
        withArrow
      >
        <ActionIcon
          variant="light"
          size="lg"
          component="label"
          htmlFor={inputId}
        >
          {pending ? <Loader size="xs" /> : <IconPlus />}
        </ActionIcon>
      </Tooltip>
      <input
        name="file"
        id={inputId}
        type="file"
        ref={inputRef}
        onChange={handleFileChange}
        hidden
      />
    </form>
  );
}
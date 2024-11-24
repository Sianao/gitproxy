import { request } from "http";
import Alert from '@mui/material/Alert';
import CheckIcon from '@mui/icons-material/Check';
import * as React from 'react';
import CustomizedSnackbars from "./Notice";
import { Button } from "@mui/material";

import { saveAs } from 'file-saver';

async function downloadFile(url: string): Promise<void> {
  try {
    url="https://gitproxy.click/"+url;
    const response = await fetch(url);

    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }

    const blob = await response.blob();

    const urlSegments = url.split('/');
    let filename = urlSegments[urlSegments.length - 1];

    saveAs(blob, filename);
  } catch (error) {
    console.error('下载文件时出错:', error);
  }
}


const GetDownloadAddress = () => {
  const [snackbarOpen, setSnackbarOpen] = React.useState(false);
  const [snackbarMessage, setSnackbarMessage] = React.useState('');
  const [snackbarSeverity, setSnackbarSeverity] = React.useState<"success" | "error">("success");

  function getDownloadAddress() {
    const req = request(
      {
        host: "jsonplaceholder.typicode.com",
        path: "/todos/1",
        method: "GET",
      },
      (response) => {
        console.log(response.statusCode);
      }
    );

    req.end();
  }

  // 更新 onClick 函数逻辑
  const onClick = (req: string) => {
    if (req == "") {
      setSnackbarSeverity("error");
      setSnackbarMessage("请输入链接再开始下载");
      setSnackbarOpen(true);
    } else if (req.startsWith('https://')) {
      setSnackbarSeverity("success");
      setSnackbarMessage("下载即将开始");
      setSnackbarOpen(true);
    } else {
      setSnackbarSeverity("error");
      setSnackbarMessage("无效的链接，请使用 HTTPS 链接。");
      setSnackbarOpen(true);
    }
    downloadFile(req)
  };

  return (
    <>
      <Button variant="contained"
        color="primary"
        size="small"
        onClick={() => {
          var address = document.getElementById("email-hero") as HTMLInputElement
          onClick(address.value)
        }}
        sx={{ minWidth: 'fit-content' }}>
        开始下载
      </Button>
      <CustomizedSnackbars
        open={snackbarOpen}
        onClose={() => setSnackbarOpen(false)}
        message={snackbarMessage}
        severity={snackbarSeverity}
      />
    </>
  );
};


export default GetDownloadAddress;


import AppAppBar from "../theme/AppAppBar"
import React from "react";

import CssBaseline from '@mui/material/CssBaseline';
import AppTheme from '../theme/AppTheme';
import Features from "../theme/Features";
import Hero from '../theme/Hero';
import Footer from "../theme/Footer";
import Head from "next/head";
export default function Home() {
  return (
    <>
      <Head>
      <title>GitHub Proxy 代理加速
      </title>
        <link rel="icon" href="favicon.ico" type="image/x-icon"/>
        <meta name='description' content='Github Proxy Github 国内代理 GitHub 文件 , Releases , archive , gist , raw.githubusercontent.com 文件代理加速下载服务.' />
        <meta name='keywords' content={`Github Proxy Github 国内代理 GitHub 文件 , Releases `} />
      
      </Head>
      <AppTheme >
        <CssBaseline enableColorScheme />
        <Hero />
        <Features></Features>
        <AppAppBar></AppAppBar>
        <Footer></Footer>
      </AppTheme>
    </>

  );
}


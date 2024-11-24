
import AppAppBar from "../theme/AppAppBar"
import React from "react";

import CssBaseline from '@mui/material/CssBaseline';
import AppTheme from '../theme/AppTheme';
import Features  from "../theme/Features";
import Hero from '../theme/Hero';
import Footer from "../theme/Footer";


export default function Home() {
  return (
    <AppTheme >
      <CssBaseline enableColorScheme />
      <Hero />
    <Features></Features>
      <AppAppBar></AppAppBar>
      <Footer></Footer>
    </AppTheme>
  );
}


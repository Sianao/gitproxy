import * as React from 'react';

import GetDownloadAddress from './logic';
import Button from '@mui/material/Button';
import Container from '@mui/material/Container';
import InputLabel from '@mui/material/InputLabel';
import Link from '@mui/material/Link';
import Stack from '@mui/material/Stack';
import TextField from '@mui/material/TextField';
import { useState } from 'react';
import Typography from '@mui/material/Typography';
import { visuallyHidden } from '@mui/utils';
import Box from '@mui/material/Box';
import SwithBar from "./SwithBar";
import Autocomplete from '@mui/material/Autocomplete';
export default function Hero() {
  const [inputValue, setInputValue] = useState(''); // 输入框的状态

  const address = [
    { label: 'https://github.com', },
  ]
  return (
    <Box
      id="hero"
      sx={(theme) => ({
        width: '100%',
        backgroundRepeat: 'no-repeat',

        backgroundImage:
          'radial-gradient(ellipse 80% 50% at 50% -20%, hsl(210, 100%, 90%), transparent)',
        ...theme.applyStyles('dark', {
          backgroundImage:
            'radial-gradient(ellipse 80% 50% at 50% -20%, hsl(210, 100%, 16%), transparent)',
        }),
      })}
    >
      <Container
        sx={{
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
          pt: { xs: 14, sm: 20 },
          pb: { xs: 8, sm: 12 },
        }}
      >
        <Stack
          spacing={2}
          useFlexGap
          sx={{ alignItems: 'center', width: { xs: '100%', sm: '70%' } }}
        >
          <Typography
            variant="h1"
            sx={{
              display: 'flex',
              flexDirection: { xs: 'column', sm: 'row' },
              alignItems: 'center',
              fontSize: 'clamp(3rem, 10vw, 3.5rem)',
            }}
          >
            Github &nbsp;
            <Typography
              component="span"
              variant="h1"
              sx={(theme) => ({
                fontSize: 'inherit',
                color: 'primary.main',
                ...theme.applyStyles('dark', {
                  color: 'primary.light',
                }),
              })}
            >
              Proxy
            </Typography>
          </Typography>
          <Typography
            sx={{
              textAlign: 'center',
              color: 'text.secondary',
              width: { sm: '100%', md: '80%' },
            }}
          >

            致力于为国内用户提供优质的 Github 代理服务
          </Typography>
          <Stack
            direction={{ xs: 'column', sm: 'row' }}
            spacing={1}
            useFlexGap
            sx={{ pt: 2, width: { xs: '100%', sm: '90%' } }}
          >
            <InputLabel htmlFor="email-hero" sx={visuallyHidden}>
              Email
            </InputLabel>
            <Autocomplete
            id="email-hero"
            freeSolo
            disableClearable
            sx={{width:"100%"}}
            options={address}
            renderInput={(params) => (
                <TextField
                    {...params}
                    hiddenLabel
                    size="small"
                    variant="outlined"
                    aria-label="Enter your GitHub address"
                    placeholder="输入您的 Github 链接 自动生成下载地址"
                    fullWidth
                   
                />
            )}
            onInputChange={(event, newInputValue) => {
                setInputValue(newInputValue);
            }}
            onChange={(event, newValue) => {
                console.log('Selected GitHub URL:', newValue);
            }}
            value={inputValue}
        />
           
            <GetDownloadAddress></GetDownloadAddress>
          </Stack>

          <Typography
            variant="caption"
            color="text.secondary"
            sx={{ textAlign: 'center' }}
          >
            点击 &quot; 开始下载 &quot; 既表示您同意 &nbsp;我们的  &nbsp;
            <Link href="#" color="primary">
              Terms & Conditions
            </Link>

          </Typography>

          <SwithBar inputValue={inputValue}></SwithBar>


        </Stack>

      </Container>
    </Box>
  );
}

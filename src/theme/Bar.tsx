import React, { useEffect, useState } from 'react';
import {
    Button,
    Dialog,
    DialogActions,
    DialogContent,
    DialogContentText,
    DialogTitle,
} from '@mui/material';

const SimpleDialogDemo = () => {
    const [open, setOpen] = useState(true);

    // 打开弹窗


    // 关闭弹窗
    const handleClose = () => {
        setOpen(false);
    };

    // 自动关闭弹窗
    // useEffect(() => {
    //     if (open) {
    //         const timer = setTimeout(() => {
    //             handleClose();
    //         }, 2000); // 3000 毫秒后关闭弹窗

    //         return () => clearTimeout(timer); 
    //     }
    // }, [open]);

    return (
        <div>
            <Dialog open={open} onClose={handleClose} sx={{ '& .MuiDialog-paper': { top: '10%', maxHeight: '80%', position: 'fixed', margin: 'auto', transition: 'all 0.3s', } }}>
                <DialogContent>
                    <DialogContentText>
                        系统添加了评论功能,如使用遇到问题,请滑动到最底部提交评论反馈。
                    </DialogContentText>
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleClose} color="primary">
                        X
                    </Button>
                </DialogActions>
            </Dialog>
        </div>
    );
};

export default SimpleDialogDemo;

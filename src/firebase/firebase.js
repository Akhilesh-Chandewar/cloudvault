import { initializeApp } from "firebase/app";
import { getAuth , GoogleAuthProvider } from "firebase/auth";

const firebaseConfig = {
    apiKey: "AIzaSyCW8VWM1DSaeofSRCoKzWDu7kTV5_sMbbs",
    authDomain: "cloud-vault-c9888.firebaseapp.com",
    projectId: "cloud-vault-c9888",
    storageBucket: "cloud-vault-c9888.appspot.com",
    messagingSenderId: "83849169528",
    appId: "1:83849169528:web:8fa1b9b4fd4338faf0fb8d",
    measurementId: "G-049X23EELR"
};

export const app = initializeApp(firebaseConfig);
export const auth = getAuth(app);
export const googleProvider = new GoogleAuthProvider();

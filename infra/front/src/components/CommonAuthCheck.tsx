import { FirebaseApp, FirebaseOptions, initializeApp } from 'firebase/app';
import { Auth, getAuth } from 'firebase/auth';

const firebaseConfig: FirebaseOptions = {
  apiKey: 'AIzaSyCt2I3KlYG2GRC2yoa0E_u7xGijnWMrBRE',
  authDomain: 'myfav-435f7.firebaseapp.com',
  projectId: 'myfav-435f7',
  storageBucket: 'myfav-435f7.appspot.com',
  messagingSenderId: '754651945016',
  appId: '1:754651945016:web:be44a5b33a4cfb58dcaecd',
  measurementId: 'G-WBYHVFX1LJ',
};

const getFireBaseApp = (): FirebaseApp => {
  return initializeApp(firebaseConfig);
};

export const getFireBaseAuth = (): Auth => {
  return getAuth(getFireBaseApp());
};

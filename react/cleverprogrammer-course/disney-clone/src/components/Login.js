import React from 'react';
import styled from 'styled-components';

function Login () {
  return (
    <Container>
      <Content>
    
        <Logo><img src="/images/disney-logos.png" alt=""/></Logo>

        <BigTitle>
            <img id="fancy-lines" src="/images/fancy-lines.png" alt=""/>
            <span>MOVIES, SHOWS,& SPORTS</span>
            <img src="/images/line.png" alt=""/>
        </BigTitle>

        <BundleButton>
          <p>Get The Disney Bundle</p>
        </BundleButton>

        <Terms></Terms>

        <Signup></Signup>

      </Content>
    </Container>
  )
}

export default Login

const Container = styled.div`
  background-color: #061435;
  position: relative;
  height: calc(100vh - 70px);

  &:before {
    position: absolute;
    content: "";
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    z-index: -1;
  }
`

const Content = styled.main`
`

const Logo = styled.div`
  display:flex;
  align-items: center;
  justify-content: center;
  padding: 100px 0px 0px;
`
const BigTitle = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 75px;
  font-weight: bold;
  width: 1950px;

  #fancy-lines {
    height: 210px;
  }
`
const BundleButton = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;
  width: 750px;
  height: 60px;
  background-color: #2157BC;
  border-radius: 4px;
`
const Terms = styled.div``
const Signup = styled.div``

import React from 'react';
import styled from 'styled-components';

function Header () {
  return (
    <Nav>
      <Logo src="/images/logo.svg"/>

      <NavMenu>
        <a>
          <img src="/images/home-icon.svg" alt="" />
          <span>HOME</span>
        </a>
        <a>
          <img src="/images/search-icon.svg" alt="" />
          <span>SEARCH</span>
        </a>
        <a>
          <img src="/images/watchlist-icon.svg" alt="" />
          <span>WATCHLIST</span>
        </a>
        <a>
          <img src="/images/original-icon.svg" alt="" />
          <span>ORIGINAL</span>
        </a>
        <a>
          <img src="/images/movie-icon.svg" alt="" />
          <span>MOVIES</span>
        </a>
        <a>
          <img src="/images/series-icon.svg" alt="" />
          <span>SERIES</span>
        </a>

      </NavMenu>
      <UserImg src="https://womenwriteaboutcomics.com/wp-content/uploads/2014/09/Screen-Shot-2014-09-11-at-20.43.04-227x200.png" />
    </Nav>
  )
}

export default Header

const Nav = styled.nav`
  height: 70px;
  background: #090b13;
  display: flex;
  align-item: center;
  padding: 0 36px;
`

const Logo = styled.img`
  width: 80px;
`

const NavMenu = styled.div`
  display: flex;
  flex: auto;
  margin-left: 125px;
  align-items: center;

  a {
    display: flex;
    align-items: center;
    padding: 0 12px;
    cursor: pointer;

    img {
      height: 20px;
    }

    span {
      font-size: 13px;
      letter-spacing: 1.42px;
    }
  }
`

const UserImg = styled.img`
  display: flex;
  align-iems: center;
  width: 40px;
  height: 40px;
  border-radius: 50%;
  cursor: pointer;
  padding: 14px 0;
`

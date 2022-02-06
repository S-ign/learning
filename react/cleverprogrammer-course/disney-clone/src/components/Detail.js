import React from 'react';
import styled from 'styled-components';

function Detail () {
  return (
    <Container>
      <Background>
        <img src="https://prod-ripcut-delivery.disney-plus.net/v1/variant/disney/0C67F69B3B1E7454EDF12614F59FDBDDFD12E34B16E0BF275D5A0EAC277FF4DF/scale?width=1440&aspectRatio=1.78&format=jpeg" alt=""/>
      </Background>

      <ImageTitle><img src="https://prod-ripcut-delivery.disney-plus.net/v1/variant/disney/DAAC1326F58FCE4E2AA847B2536F64D665B7598D1BDD26E4D937BA6839F2DC38/scale?width=1440&aspectRatio=1.78&format=png" alt=""/></ImageTitle>

      <SubTitleInfo>
        <Icons>
          <img src="https://prod-ripcut-delivery.disney-plus.net/v1/variant/disney/7843B7AF26846E842BA3EDD8628471E6435C5DAC9F89FF5307D345EAADCAE6F5/scale?width=240" alt=""/>
          <img src="https://prod-ripcut-delivery.disney-plus.net/v1/variant/disney/FD4912EB883B7CCB847EB9C62E1FC853D547CAF7DF940B9086AE35AFAD0848AB/scale?width=240" alt=""/>
          <img src="https://prod-ripcut-delivery.disney-plus.net/v1/variant/disney/FAE63AC7AC11C27C949E1856CF188BF09FC20EA52AEA3B65B43C24EEB5F29BFD/scale?width=240" alt=""/>
          <span>
            1992
            <ul>
              <li>1h 34m</li>
            </ul>
          </span>
        </Icons>

        <Catagory>
          <span>Family, Fantasy, Animation, Action-Adventure, Musical</span>
        </Catagory>
      </SubTitleInfo>

      <Controls>

        <PlayButton>
          <img src="/images/play-icon-black.png" alt=""/>
          <span>PLAY</span>
        </PlayButton>

        <TrailerButton>
          <img src="/images/play-icon-white.png" alt=""/>
          <span>Trailer</span>
        </TrailerButton>

        <AddButton>
          <span>+</span>
        </AddButton>

        <GroupWatchButton>
          <img src="/images/group-icon.png" alt=""/>
        </GroupWatchButton>
      </Controls>

      <Description>
        <p>Street-smart Aladdin and Princess Jasmine join forces to save the kingdom from the evil sorcerer Jafar. Along the way, Aladdin learns to believe in himself with help from a hilarious Genie...and three wishes that can change everything.</p>
      </Description>
      
    </Container>
  )
}

export default Detail

const Container = styled.div`
  min-height: calc(100vh - 70px);
  padding: 0 calc(3.5vw + 5px);
  position: relative;
`

const Background = styled.div`
  position: fixed;
  top: 0;
  left: 0;
  bottom: 0;
  right: 0;
  z-index: -1;
  opacity: 0.8;

  img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }

`

const ImageTitle = styled.div`
  height: 30vh;
  min-height: 170px;
  min-width: 200px;
  width: 35vw;

  img {
    height: 100%;
    width: 100%;
    object-fit: contain;
  }
  
`
const SubTitleInfo = styled.div`
`
const Icons = styled.div`
  display: flex;
  align-items: center;

  span {
    display:flex;
    align-items: center;
    padding: 0;
    margin: 0;
  }

  ul {
    padding: 16px;
    margin: 0;
  }

  img {
    height: 30px;
    width: 30px;
    padding: 5px; 5px;
  }
`
const Catagory = styled.div`
  font-size: 12px;
  margin-bottom: 32px;
`

const Controls = styled.div`
  display: flex;
  align-items: center;
`

const PlayButton = styled.button`
  border-radius: 4px;
  font-size: 15px;
  display: flex;
  align-items: center;
  height: 56px;
  background: rgb(249, 249, 249);
  border: none;
  padding: 0px 24px;
  margin-right: 22px;
  letter-spacing: 1.8px;
  curser: pointer;

  &:hover {
    background: rgb(198, 198, 198)
  }
`
const TrailerButton = styled(PlayButton)`
  background: rgba(0, 0, 0, 0.3);
  border: 1px solid rgb(249, 249, 249);
  color: rgb(249, 249, 249);
  text-transform: uppercase;
`
const AddButton = styled.button`
  width: 44px;
  height: 44px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  border: 2px solid white;
  background-color: rgba(0, 0, 0, 0.6);
  cursor: pointer;
  margin-right: 16px;
  
  span {
    font-size: 30px;
    color: white;
  }
`
const GroupWatchButton = styled(AddButton)`
  background: rgb(0, 0, 0);
`

const Description = styled.div`
  width: 1000px;
  font-size: 23px;
`

import React from "react";
import styles from "./homePage.module.css";
import Slider from "@/components/slider/Slider";
import Image from "next/image";
import Where from "@/components/Where/Where";
import Link from "next/link";
import Advertise from "@/components/advertiseTravel/Advertise";
import { Button } from "react-bootstrap";


const HomePage = () => {

  return (
    <div>
      <Slider/>
      
      <div className={styles.commentSection}>
        <div className={styles.ingredient}>
          <h3>Uygulamayı kullananlardan okuyun !</h3>
        </div>

        <section className={styles.products}>
          <div className={styles.boxContainer}>
            <div className={styles.box}>
              <div className={styles.boxHead}>
                <div className={styles.userInfo}>
                  <Image alt="user" src="/navbarLogo.png" width={32} height={32}/>
                  <span className={styles.commentUser}>Kullanıcı</span>
                </div>
                <p className={styles.name}>
                  Bu uygulamayı kullanarak çok iyi arkadaşlar edindim. <br />{" "}
                  Kesinlikle bu uygulamayı tavsiye ediyorum. Hem güvenli,
                  hem kullanışlı.
                </p>
              </div>
            </div>
            <div className={styles.box}>
              <div className={styles.boxHead}>
              <div className={styles.userInfo}>
                  <Image alt="user" src="/navbarLogo.png" width={32} height={32}/>
                  <span className={styles.commentUser}>Kullanıcı</span>
                </div>
                <p className={styles.name}>
                  Bu uygulamayı kullanarak çok iyi arkadaşlar edindim.
                  Kesinlikle bu uygulamayı tavsiye ediyorum.Hem güvenli, hem
                  kullanışlı.
                </p>
              </div>
            </div>
            <div className={styles.box}>
              <div className={styles.boxHead}>
              <div className={styles.userInfo}>
                  <Image alt="user" src="/navbarLogo.png" width={32} height={32}/>
                  <span className={styles.commentUser}>Kullanıcı</span>
                </div>
                <p className={styles.name}>
                  Bu uygulamayı kullanarak çok iyi arkadaşlar edindim. <br />{" "}
                  Kesinlikle bu uygulamayı tavsiye ediyorum. Hem güvenli,
                  hem kullanışlı.
                </p>
              </div>
            </div>
          </div>
        </section>
      </div>
      <div className={styles.ingredient}>
        <h3>Neden Yol Arkadaşımı seçmelisin ?</h3>
      </div>

      <div className={styles.buto}>
        <div className={styles.icon}>
          <Image alt="safe" src="/safe.png" width={32} height={32}/>
          <p>Güvenilir</p>
        </div>
        <div className={styles.icon}>
          <Image alt="safe" src="/friendly.png" width={32} height={32}/>
          <p>Kullanıcı Dostu</p>
        </div>
        <div className={styles.icon}>
          <Image alt="safe" src="/fast.png" width={32} height={32}/>
          <p>Hızlı</p>
        </div>
      </div>


      <div className={styles.findAdvertSection}>
      <div className={styles.whereTitle}>
        <h3>Nereye gitmek istersiniz ?</h3>
      </div>

      <div className={styles.findAdvert}>
        <div className={styles.city}>
              <Where type="from"/>
             <Where type="to"/>
        </div>
        <div className={styles.button}>
          <Link href="/advert" className={styles.buttonLink}>
          <Button variant="primary">
            İlan ara
          </Button>
          </Link>
        </div>
      </div>
      </div>

      <div className={styles.toPublish}>
      <div className={styles.findAdvertTitle}>
        <h3>Aradığın yolculuk yok mu?</h3>
      </div>
      <div className={styles.createAdvert}>
        <p className={styles.name}>
          Aradığın yolculuk ilanı yoksa ve yol arkadaşı arıyorsan hemen bir ilan
          paylaş!
        </p>
        <div className={styles.button}>
          <Advertise className={styles.advertPublish}/>
        </div>
      </div>
      </div>
    </div>
  );
};
export default HomePage;

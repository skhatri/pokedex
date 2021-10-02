package com.api;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties(ignoreUnknown = true)
public class Pokemon {
    @JsonProperty("s_no")
    private int pokedexNumber;
    private String name;
    private int generation;
    private String status;
    private String species;
    private int typeNumber;
    private String type1;
    private String type2;
    private float height;
    private float weight;
    private int abilities;
    private String ability1;
    private String ability2;
    @JsonProperty("ability_hidden")
    private String hiddenAbility;
    private String totalPoints;
    private int hp;
    private int attack;
    private int defense;
    private int spAttack;
    private int spDefense;
    private int speed;
    private float catchRate;
    private String baseFriendship;
    private int baseExperience;
    private int growthRate;
    private int eggTypeNumber;
    private int eggType1;
    private int eggType2;
    private int percentageMale;
    private int eggCycles;
    private float againstNormal;
    private float againstFire;
    private float againstWater;
    private float againstElectric;
    private float againstGrass;
    private float AgainstIce;
    private float AgainstFight;
    private float againstPoison;
    private float againstGround;
    private float againstFlying;
    private float againstPsychic;
    private float againstBug;
    private float againstRock;
    private float againstGhost;
    private float againstDragon;
    private float againstDark;
    private float againstSteel;
    private float againstFairy;

    public int getPokedexNumber() {
        return pokedexNumber;
    }

    public void setPokedexNumber(int pokedexNumber) {
        this.pokedexNumber = pokedexNumber;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public int getGeneration() {
        return generation;
    }

    public void setGeneration(int generation) {
        this.generation = generation;
    }

    public String getStatus() {
        return status;
    }

    public void setStatus(String status) {
        this.status = status;
    }

    public String getSpecies() {
        return species;
    }

    public void setSpecies(String species) {
        this.species = species;
    }

    public int getTypeNumber() {
        return typeNumber;
    }

    public void setTypeNumber(int typeNumber) {
        this.typeNumber = typeNumber;
    }

    public String getType1() {
        return type1;
    }

    public void setType1(String type1) {
        this.type1 = type1;
    }

    public String getType2() {
        return type2;
    }

    public void setType2(String type2) {
        this.type2 = type2;
    }

    public float getHeight() {
        return height;
    }

    public void setHeight(float height) {
        this.height = height;
    }

    public float getWeight() {
        return weight;
    }

    public void setWeight(float weight) {
        this.weight = weight;
    }

    public int getAbilities() {
        return abilities;
    }

    public void setAbilities(int abilities) {
        this.abilities = abilities;
    }

    public String getAbility1() {
        return ability1;
    }

    public void setAbility1(String ability1) {
        this.ability1 = ability1;
    }

    public String getAbility2() {
        return ability2;
    }

    public void setAbility2(String ability2) {
        this.ability2 = ability2;
    }

    public String getHiddenAbility() {
        return hiddenAbility;
    }

    public void setHiddenAbility(String hiddenAbility) {
        this.hiddenAbility = hiddenAbility;
    }

    public String getTotalPoints() {
        return totalPoints;
    }

    public void setTotalPoints(String totalPoints) {
        this.totalPoints = totalPoints;
    }

    public int getHp() {
        return hp;
    }

    public void setHp(int hp) {
        this.hp = hp;
    }

    public int getAttack() {
        return attack;
    }

    public void setAttack(int attack) {
        this.attack = attack;
    }

    public int getDefense() {
        return defense;
    }

    public void setDefense(int defense) {
        this.defense = defense;
    }

    public int getSpAttack() {
        return spAttack;
    }

    public void setSpAttack(int spAttack) {
        this.spAttack = spAttack;
    }

    public int getSpDefense() {
        return spDefense;
    }

    public void setSpDefense(int spDefense) {
        this.spDefense = spDefense;
    }

    public int getSpeed() {
        return speed;
    }

    public void setSpeed(int speed) {
        this.speed = speed;
    }

    public float getCatchRate() {
        return catchRate;
    }

    public void setCatchRate(float catchRate) {
        this.catchRate = catchRate;
    }

    public String getBaseFriendship() {
        return baseFriendship;
    }

    public void setBaseFriendship(String baseFriendship) {
        this.baseFriendship = baseFriendship;
    }

    public int getBaseExperience() {
        return baseExperience;
    }

    public void setBaseExperience(int baseExperience) {
        this.baseExperience = baseExperience;
    }

    public int getGrowthRate() {
        return growthRate;
    }

    public void setGrowthRate(int growthRate) {
        this.growthRate = growthRate;
    }

    public int getEggTypeNumber() {
        return eggTypeNumber;
    }

    public void setEggTypeNumber(int eggTypeNumber) {
        this.eggTypeNumber = eggTypeNumber;
    }

    public int getEggType1() {
        return eggType1;
    }

    public void setEggType1(int eggType1) {
        this.eggType1 = eggType1;
    }

    public int getEggType2() {
        return eggType2;
    }

    public void setEggType2(int eggType2) {
        this.eggType2 = eggType2;
    }

    public int getPercentageMale() {
        return percentageMale;
    }

    public void setPercentageMale(int percentageMale) {
        this.percentageMale = percentageMale;
    }

    public int getEggCycles() {
        return eggCycles;
    }

    public void setEggCycles(int eggCycles) {
        this.eggCycles = eggCycles;
    }

    public float getAgainstNormal() {
        return againstNormal;
    }

    public void setAgainstNormal(float againstNormal) {
        this.againstNormal = againstNormal;
    }

    public float getAgainstFire() {
        return againstFire;
    }

    public void setAgainstFire(float againstFire) {
        this.againstFire = againstFire;
    }

    public float getAgainstWater() {
        return againstWater;
    }

    public void setAgainstWater(float againstWater) {
        this.againstWater = againstWater;
    }

    public float getAgainstElectric() {
        return againstElectric;
    }

    public void setAgainstElectric(float againstElectric) {
        this.againstElectric = againstElectric;
    }

    public float getAgainstGrass() {
        return againstGrass;
    }

    public void setAgainstGrass(float againstGrass) {
        this.againstGrass = againstGrass;
    }

    public float getAgainstIce() {
        return AgainstIce;
    }

    public void setAgainstIce(float againstIce) {
        AgainstIce = againstIce;
    }

    public float getAgainstFight() {
        return AgainstFight;
    }

    public void setAgainstFight(float againstFight) {
        AgainstFight = againstFight;
    }

    public float getAgainstPoison() {
        return againstPoison;
    }

    public void setAgainstPoison(float againstPoison) {
        this.againstPoison = againstPoison;
    }

    public float getAgainstGround() {
        return againstGround;
    }

    public void setAgainstGround(float againstGround) {
        this.againstGround = againstGround;
    }

    public float getAgainstFlying() {
        return againstFlying;
    }

    public void setAgainstFlying(float againstFlying) {
        this.againstFlying = againstFlying;
    }

    public float getAgainstPsychic() {
        return againstPsychic;
    }

    public void setAgainstPsychic(float againstPsychic) {
        this.againstPsychic = againstPsychic;
    }

    public float getAgainstBug() {
        return againstBug;
    }

    public void setAgainstBug(float againstBug) {
        this.againstBug = againstBug;
    }

    public float getAgainstRock() {
        return againstRock;
    }

    public void setAgainstRock(float againstRock) {
        this.againstRock = againstRock;
    }

    public float getAgainstGhost() {
        return againstGhost;
    }

    public void setAgainstGhost(float againstGhost) {
        this.againstGhost = againstGhost;
    }

    public float getAgainstDragon() {
        return againstDragon;
    }

    public void setAgainstDragon(float againstDragon) {
        this.againstDragon = againstDragon;
    }

    public float getAgainstDark() {
        return againstDark;
    }

    public void setAgainstDark(float againstDark) {
        this.againstDark = againstDark;
    }

    public float getAgainstSteel() {
        return againstSteel;
    }

    public void setAgainstSteel(float againstSteel) {
        this.againstSteel = againstSteel;
    }

    public float getAgainstFairy() {
        return againstFairy;
    }

    public void setAgainstFairy(float againstFairy) {
        this.againstFairy = againstFairy;
    }
}

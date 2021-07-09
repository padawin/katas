def update(item):
    if item["sellIn"] <= 0:
        qualityEvolution = -item["quality"]
    elif item["sellIn"] <= 5:
        qualityEvolution = 3
    elif item["sellIn"] <= 10:
        qualityEvolution = 2
    else:
        qualityEvolution = 1
    return qualityEvolution, -1

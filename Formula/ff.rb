class Ff < Formula
  desc "A brief description of your application"
  homepage "https://github.com/joshwycuff/homebrew-toolbox"
  url "https://github.com/joshwycuff/homebrew-toolbox/releases/download/ff-f374b4c95ad2cc9913228b2927d60be0823c7304/ff.tar.gz"
  version "f374b4c95ad2cc9913228b2927d60be0823c7304"

  depends_on "rg"
  depends_on "bat"

  def install
    bin.install "ff"
    man1.install Dir["*.1"]
    bash_completion.install "ff.bash" => "ff"
    zsh_completion.install "ff.zsh" => "_ff"
    fish_completion.install "ff.fish"
  end
end

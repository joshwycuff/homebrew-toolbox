class Fman < Formula
  desc "Perform fuzzy search on man pages."
  homepage "https://github.com/joshwycuff/homebrew-toolbox"
  url "https://github.com/joshwycuff/homebrew-toolbox/releases/download/fman-b185f0ccd960959ee20ad2c2288a06d01da8119/fman",
  version "b185f0ccd960959ee20ad2c2288a06d01da8119"

  depends_on "apropos"
  depends_on "fzf"

  def install
    bin.install "tools/fman/fman"
  end
end
